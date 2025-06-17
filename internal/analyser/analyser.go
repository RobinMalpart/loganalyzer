package analyzer

import (
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

func RunAll(logs []config.LogConfig) []Result {
	var wg sync.WaitGroup
	resultChan := make(chan Result, len(logs))

	rand.Seed(time.Now().UnixNano())

	for _, logCfg := range logs {
		wg.Add(1)
		go func(log config.LogConfig) {
			defer wg.Done()
			res := analyzeLog(log)
			resultChan <- res
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

func analyzeLog(log config.LogConfig) Result {
	var res Result
	res.LogID = log.ID
	res.FilePath = log.Path

	_, err := os.Stat(log.Path)
	if err != nil {
		if os.IsNotExist(err) {
			res.Status = "FAILED"
			res.Message = "Fichier introuvable."
			res.ErrorDetails = ErrFileNotFound{Path: log.Path}.Error()
			return res
		}
		res.Status = "FAILED"
		res.Message = "Erreur d'accès au fichier."
		res.ErrorDetails = err.Error()
		return res
	}

	sleepMs := rand.Intn(151) + 50
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	if rand.Float64() < 0.1 {
		parseErr := ErrParsing{Path: log.Path}
		res.Status = "FAILED"
		res.Message = "Erreur de parsing."
		res.ErrorDetails = parseErr.Error()
		return res
	}

	res.Status = "OK"
	res.Message = "Analyse terminée avec succès."
	res.ErrorDetails = ""
	return res
}
