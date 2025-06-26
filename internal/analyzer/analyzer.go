package analyzer

import (
	"errors"
	"loganalyzer/internal/config"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Result struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

func Analyse(logs []config.LogConfig) []Result {
	var wg sync.WaitGroup
	resultChan := make(chan Result, len(logs))

	for _, logCfg := range logs {
		wg.Add(1)
		go func(log config.LogConfig) {
			defer wg.Done()
			err := analyzeLog(log.Path)
			message, status := getMessageAndStatus(err)
			result := Result{
				LogID:    log.ID,
				FilePath: log.Path,
				Status:   status,
				Message:  message,
			}
			if err != nil {
				result.ErrorDetails = err.Error()
			}
			resultChan <- result
		}(logCfg)
	}

	wg.Wait()
	close(resultChan)

	var results []Result
	for res := range resultChan {
		results = append(results, res)
	}

	return results
}

func analyzeLog(logPath string) error {
	if _, err := os.Stat(logPath); err != nil {
		return &ErrFileNotFound{Path: logPath, Err: ErrFileNotFoundSentinel}
	}

	sleepMs := rand.Intn(151) + 50
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	content, err := os.ReadFile(logPath)
	if err != nil {
		return &ErrFileNotFound{Path: logPath, Err: ErrFileNotFoundSentinel}
	}

	if len(content) == 0 {
		return &ErrFileEmpty{Path: logPath, Err: ErrFileEmptySentinel}
	}
	return nil
}

func getMessageAndStatus(err error) (string, string) {
	if err == nil {
		return "Analyse terminée avec succès.", "OK"
	}

	if errors.Is(err, ErrFileNotFoundSentinel) {
		return "Fichier introuvable.", "FAILED"
	}

	if errors.Is(err, ErrFileEmptySentinel) {
		return "Fichier vide.", "FAILED"
	}

	return "Erreur inconnue.", "FAILED"
}
