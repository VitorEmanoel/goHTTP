package goHTTP

import "strings"

func clearArray(array []string) []string{
	var finalArray [] string
	for i:= 0; i < len(array); i++{
		if array[i] != "" && array[i] != " "{
			finalArray = append(finalArray, array[i])
		}
	}
	return finalArray
}

func isRouter(url string, path string) bool{
	urlSplited := clearArray(strings.Split(url, "/"))
	pathSplited := clearArray(strings.Split(path, "/"))
	if len(pathSplited) > len(urlSplited){
		return false
	}
	if len(pathSplited) == len(urlSplited){
		for i := 0; i < len(pathSplited); i++{
			if pathSplited[i] == urlSplited[i] || pathSplited[i][0] == ':'{
				return true
			}
			return false
		}
	}
	return false
}
