// Author: Bruno Lucena <bvlg900f@gmail.com>
package handler

type Response struct {
	Success  bool
	Errors   []string
	Messages []string
	Result   []Result
}

type Result struct {
	Id   string
	Name string
}
