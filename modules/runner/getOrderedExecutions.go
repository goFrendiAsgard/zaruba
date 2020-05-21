package runner

import (
	"sort"

	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

func getOrderedExecutions(p *config.ProjectConfig, executedServices []string) (executions [][]string, err error) {
	services := includeCandidate(p, executedServices)
	executions = [][]string{}
	leftServices := getLeftServices(services, executions)
	for len(leftServices) > 0 {
		currentBatch := []string{}
		for _, service := range leftServices {
			component, err := p.GetComponentByName(service)
			if err != nil {
				return executions, err
			}
			if isDependenciesFullfiled(component, executions) {
				currentBatch = append(currentBatch, service)
			}
		}
		sort.Strings(currentBatch)
		executions = append(executions, currentBatch)
		leftServices = getLeftServices(services, executions)
	}
	return executions, err
}

func includeCandidate(p *config.ProjectConfig, services []string) []string {
	dependencies := []string{}
	for _, service := range services {
		component, err := p.GetComponentByName(service)
		if err != nil {
			logger.Fatal("Cannot get component ", service)
		}
		componentDependencies := component.GetDependencies()
		completeDependencies := includeCandidate(p, componentDependencies)
		for _, dependency := range completeDependencies {
			if !inArray(dependency, services) && !inArray(dependency, dependencies) {
				dependencies = append(dependencies, dependency)
			}
		}
	}
	return append(dependencies, services...)
}

func getLeftServices(services []string, executions [][]string) (left []string) {
	flattenExecutions := flattenArray(executions)
	left = []string{}
	for _, service := range services {
		if !inArray(service, flattenExecutions) {
			left = append(left, service)
		}
	}
	return left
}

func inArray(element string, arr []string) (found bool) {
	found = false
	for _, otherElement := range arr {
		if element == otherElement {
			found = true
			break
		}
	}
	return found
}

func isDependenciesFullfiled(component *config.Component, executions [][]string) (fullfilled bool) {
	flattenExecutions := flattenArray(executions)
	for _, dependency := range component.GetDependencies() {
		found := false
		for _, service := range flattenExecutions {
			if service == dependency {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func flattenArray(arr [][]string) (flatten []string) {
	flatten = []string{}
	for _, subArr := range arr {
		for _, element := range subArr {
			flatten = append(flatten, element)
		}
	}
	return flatten
}
