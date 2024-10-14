package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type QuizRequest struct {
    Topic       string `json:"topic"`
    NumQuestions int    `json:"num_questions"`
}

func fetchQuiz(topic string, numQuestions int) (string, error) {
    requestBody, err := json.Marshal(QuizRequest{
        Topic:       topic,
        NumQuestions: numQuestions,
    })
    if err != nil {
        return "", err
    }

    resp, err := http.Post("http://127.0.0.1:8000/get-quiz", "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)

    quiz, ok := result["quiz"].(string)
    if !ok {
        return "", fmt.Errorf("invalid response")
    }

    return quiz, nil
}
