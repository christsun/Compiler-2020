package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// /fetchTask
func getTask(w http.ResponseWriter, r *http.Request) {
	if len(semanticPool) == 0 && len(codegenPool) == 0 && len(optimizePool) == 0{
		_ = json.NewEncoder(w).Encode(simpleSendFormat{
			Code:    404,
			Message: "No available work currently",
		})
		return
	}
	if len(semanticPool) != 0 {
		poolElement := semanticPool[0]
		var runningList []string
		remainPend := len(poolElement.pending)
		var idx = 0
		for ; remainPend == 0 && idx < len(semanticPool); idx++ {
			poolElement = semanticPool[idx]
			remainPend = len(poolElement.pending)
			if remainPend != 0 {
				break
			}
		}
		if remainPend != 0 {
			if remainPend < 5 {
				runningList = make([]string, remainPend)
				_ = copy(runningList, semanticPool[idx].pending)
				semanticPool[idx].pending = append(semanticPool[idx].pending[remainPend:])
				for _, data := range semanticPool[idx].pending {
					semanticPool[idx].runningSet[data] = true
				}
			} else {
				runningList = make([]string, 5)
				_ = copy(runningList, semanticPool[idx].pending[0:5])
				semanticPool[idx].pending = append(semanticPool[idx].pending[5:])
				for idx2, data := range semanticPool[idx].pending {
					semanticPool[idx].runningSet[data] = true
					if idx2 == 4 {
						break
					}
				}
			}
			sqlCmd := fmt.Sprintf("SELECT sema_uid, sema_sourceCode, sema_assertion, sema_timeLimit, sema_memoryLimit FROM dataset_semantic WHERE %s", "sema_uid='"+strings.Join(runningList, "' OR sema_uid='")+"'")
			result, err := executionQuery(sqlCmd)
			if err != nil {
				logger(fmt.Sprintf("Runtime error: %s", err.Error()), 1)
				_ = json.NewEncoder(w).Encode(simpleSendFormat{
					Code:    400,
					Message: fmt.Sprint(err.Error()),
				})
				return
			}
			var sentReq requestSemanticTaskFormat
			sentReq.Code = 200
			sentReq.Target = make([]subtaskSemanticFormat, 0, 5)
			for result.Next() {
				var id string
				var sourceCode string
				var assert string
				var timeLimit float32
				var memoryLimit int
				err = result.Scan(&id, &sourceCode, &assert, &timeLimit, &memoryLimit)
				if err != nil {
					logger(fmt.Sprintf("SQL Runtime error: %s", err.Error()), 1)
				}
				sentReq.Target = append(sentReq.Target, subtaskSemanticFormat{
					Uuid:            poolElement.uuid,
					Repo:            poolElement.repo,
					TestCase:        id,
					Stage:           1,
					Subworkid:       id + "_" + n.Next(),
					InputSourceCode: sourceCode,
					Assertion:       assert,
					TimeLimit:       timeLimit,
					MemoryLimit:     memoryLimit,
					TaskID:          poolElement.recordID,
				})
			}

			err = json.NewEncoder(w).Encode(sentReq)
			if err != nil {
				logger(fmt.Sprintf("Runtime error: %s", err.Error()), 1)
			}
			return
		}
	}
	if len(codegenPool) != 0 {
		poolElement := codegenPool[0]
		var runningList []string
		remainPend := len(poolElement.pending)
		var idx = 0
		for ; remainPend == 0 && idx < len(codegenPool); idx++ {
			poolElement = codegenPool[idx]
			remainPend = len(poolElement.pending)
			if remainPend != 0 {
				break
			}
		}
		if remainPend != 0 {
			if remainPend < 5 {
				codegenPool[idx].running = make([]string, remainPend)
				runningList = make([]string, remainPend)
				_ = copy(codegenPool[idx].running, codegenPool[idx].pending)
				codegenPool[idx].pending = append(codegenPool[idx].pending[remainPend:])
			} else {
				codegenPool[idx].running = make([]string, 5)
				runningList = make([]string, 5)
				_ = copy(codegenPool[idx].running, codegenPool[idx].pending[0:5])
				codegenPool[idx].pending = append(codegenPool[idx].pending[5:])
			}
			_ = copy(runningList, codegenPool[idx].running)
			cmd := "SELECT cg_uid, cg_sourceCode, cg_assertion, cg_timeLimit, cg_memoryLimit, cg_inputCtx, cg_outputCtx, cg_outputCode FROM dataset_codegen WHERE " +
				"cg_uid='" + strings.Join(runningList, "' OR cg_uid='") + "'"
			fmt.Printf("Execution Sentence:%s\n", cmd)
			result, err := executionQuery(cmd)
			if err != nil {
				fmt.Printf("Error %s", err.Error())
				return
			}
			var sentReq requestCodegenTaskFormat
			sentReq.Code = 200
			sentReq.Target = make([]subtaskCodegenFormat, 0, 5)
			for result.Next() {
				var id string
				var sourceCode string
				var assert string
				var timeLimit float32
				var memoryLimit int
				var inputContext string
				var outputContext string
				var outputCode int
				err = result.Scan(&id, &sourceCode, &assert, &timeLimit, &memoryLimit, &inputContext, &outputContext, &outputCode)
				if err != nil {
					fmt.Printf("runtime warning:%s when scanning the codegen database", err.Error())
				}
				sentReq.Target = append(sentReq.Target, subtaskCodegenFormat{
					Uuid:            poolElement.uuid,
					Repo:            poolElement.repo,
					TestCase:        id,
					Stage:           2,
					Subworkid:       id + "_" + n.Next(),
					InputSourceCode: sourceCode,
					InputContent:    inputContext,
					OutputCode:      outputCode,
					OutputContent:   outputContext,
					TimeLimit:       timeLimit,
					MemoryLimit:     memoryLimit,
					TaskID:          poolElement.recordID,
				})
			}

			err = json.NewEncoder(w).Encode(sentReq)
			if err != nil {
				fmt.Printf("runtime error: %s", err.Error())
			}
			return
		}
	}
	if len(optimizePool) != 0 {
		poolElement := optimizePool[0]
		var runningList []string
		remainPend := len(poolElement.pending)
		var idx = 0
		for ; remainPend == 0 && idx < len(optimizePool); idx++ {
			poolElement = optimizePool[idx]
			remainPend = len(poolElement.pending)
			if remainPend != 0 {
				break
			}
		}
		if remainPend != 0 {
			if remainPend < 5 {
				optimizePool[idx].running = make([]string, remainPend)
				runningList = make([]string, remainPend)
				_ = copy(optimizePool[idx].running, optimizePool[idx].pending)
				optimizePool[idx].pending = append(optimizePool[idx].pending[remainPend:])
			} else {
				optimizePool[idx].running = make([]string, 5)
				runningList = make([]string, 5)
				_ = copy(optimizePool[idx].running, optimizePool[idx].pending[0:5])
				optimizePool[idx].pending = append(optimizePool[idx].pending[5:])
			}
			_ = copy(runningList, optimizePool[idx].running)
			cmd := "SELECT optim_uid, optim_sourceCode, optim_assertion, optim_timeLimit, optim_memoryLimit, optim_inputCtx, optim_outputCtx, optim_outputCode FROM dataset_optimize WHERE " +
				"optim_uid='" + strings.Join(runningList, "' OR optim_uid='") + "'"
			fmt.Printf("Execution Sentence:%s\n", cmd)
			result, err := executionQuery(cmd)
			if err != nil {
				fmt.Printf("Error %s", err.Error())
				return
			}
			var sentReq requestCodegenTaskFormat
			sentReq.Code = 200
			sentReq.Target = make([]subtaskCodegenFormat, 0, 5)
			for result.Next() {
				var id string
				var sourceCode string
				var assert string
				var timeLimit float32
				var memoryLimit int
				var inputContext string
				var outputContext string
				var outputCode int
				err = result.Scan(&id, &sourceCode, &assert, &timeLimit, &memoryLimit, &inputContext, &outputContext, &outputCode)
				if err != nil {
					fmt.Printf("runtime warning:%s when scanning the codegen database", err.Error())
				}
				sentReq.Target = append(sentReq.Target, subtaskCodegenFormat{
					Uuid:            poolElement.uuid,
					Repo:            poolElement.repo,
					TestCase:        id,
					Stage:           3,
					Subworkid:       id + "_" + n.Next(),
					InputSourceCode: sourceCode,
					InputContent:    inputContext,
					OutputCode:      outputCode,
					OutputContent:   outputContext,
					TimeLimit:       timeLimit,
					MemoryLimit:     memoryLimit,
					TaskID:          poolElement.recordID,
				})
			}

			err = json.NewEncoder(w).Encode(sentReq)
			if err != nil {
				fmt.Printf("runtime error: %s", err.Error())
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(requestSemanticTaskFormat{
		Code:   1,
		Target: nil,
	})
	if err != nil {
		fmt.Printf("runtime error: %s", err.Error())
	}
	return
}
