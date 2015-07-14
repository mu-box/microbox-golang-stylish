// Copyright (c) 2015 Pagoda Box Inc
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v.
// 2.0. If a copy of the MPL was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.
//

package stylish

import (
  "fmt"
  "strings"

  "github.com/mitchellh/go-wordwrap"
)

// Header styles and prints a header as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#header
//
// Usage:
// Header "i am a header"
//
// Output:
// :::::::::::::::::::::::::: I AM A HEADER :::::::::::::::::::::::::
func Header(header string) string {

  maxLen := 70
  subLen := len(fmt.Sprintf("%v", header))

  leftLen := (maxLen - subLen)/2 + (maxLen - subLen)%2
  rightLen := (maxLen - subLen)/2

  // print the header, inserting a ':' (colon) 'n' times, where 'n' is the number
  // remaining after subtracting subLen (number of 'reserved' characters) from
  // maxLen (maximum number of allowed characters)
  return fmt.Sprintf(`
%v
`, fmt.Sprintf("%v %v %v", strings.Repeat(":", leftLen), strings.ToUpper(header), strings.Repeat(":", rightLen)))
}

// ProcessStart styles and prints a 'child process' as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#child-process
//
// Usage:
// ProcessStart "i am a process"
//
// Output:
// I AM A PROCESS :::::::::::::::::::::::::::::::::::::::::::::::: =>
func ProcessStart(process string) string {

  maxLen := 70
  subLen := len(fmt.Sprintf("%v=>", process))

  // print the process, inserting a ':' (colon) 'n' times, where 'n' is the number
  // remaining after subtracting subLen (number of 'reserved' characters) from
  // maxLen (maximum number of allowed characters)
  return fmt.Sprintf(`
%v
`, fmt.Sprintf("%v %v =>", strings.ToUpper(process), strings.Repeat(":", (maxLen-subLen))))
}

// ProcessEnd styles and prints a 'child process' as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#child-process
//
// Usage:
// ProcessEnd "i am a process"
//
// Output:
// <= :::::::::::::::::::::::::::::::::::::::::::: END I AM A PROCESS
func ProcessEnd(process string) string {

  maxLen := 70
  subLen := len(fmt.Sprintf("<=%v[√]", process))

  // print the process, inserting a ':' (colon) 'n' times, where 'n' is the number
  // remaining after subtracting subLen (number of 'reserved' characters) from
  // maxLen (maximum number of allowed characters)
  return fmt.Sprintf(`
%v
`, fmt.Sprintf("<= %v %v [√]", strings.Repeat(":", (maxLen-subLen)), strings.ToUpper(process)))
}

// SubTask styles and prints a 'sub task' as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#sub-tasks
//
// Usage:
// SubTask "i am a sub task"
//
// Output:
// ::::::::: I AM A SUB TASK
func SubTask(task string) string {
  return fmt.Sprintf(`
::::::::: %v
`, strings.ToUpper(task))
}

// SubTaskSuccess styles and prints a footer to a successful subtask
//
// Usage:
// SubTaskSuccess
//
// Output:
// <<<<<<<<< [√] SUCCESS
func SubTaskSuccess() string {
  return fmt.Sprintf("\n<<<<<<<<< [√] SUCCESS\n")
}

// SubTaskFail styles and prints a footer to a failed subtask
//
// Usage:
// SubTaskFail
//
// Output:
// <<<<<<<<< [!] FAILED
func SubTaskFail() string {
  return fmt.Sprintf("\n<<<<<<<<< [!] FAILED\n")
}

// Bullet styles and prints a message as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#bullet-points
//
// Usage:
// Bullet "i am a bullet"
// Bullet []string{"we", "are", "many", "bullets"}
//
// Output:
// +> i am a bullet
//
// +> we
// +> are
// +> many
// +> bullets
func Bullet(v interface{}) string {

  // if it's a slice of strings, iterate over each one printing them individually...
  switch t := v.(type) {
  case []string:
    for i := range t {
      return fmt.Sprintf("+> %v\n", t[i])
    }
  }

  // otherwise just print the value
  return fmt.Sprintf("+> %v\n", v)
}

// Warning styles and prints a message as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#warning
//
// Usage:
// Warning "You just bought Hot Pockets!"
//
// Output:
// -----------------------------  WARNING  -----------------------------
// You just bought Hot Pockets!
func Warning(body string) string {
  return fmt.Sprintf(`
-----------------------------  WARNING  -----------------------------
%v
`, wordwrap.WrapString(body, 70))
}

// Error styles and prints a message as outlined at:
// http://nanodocs.gopagoda.io/engines/style-guide#fatal_errors
//
// Usage:
// Error "nuclear launch detected", "All your base are belong to us"
//
// Output:
// ! NUCLEAR LAUNCH DETECTED !
//
// All your base are belong to us
func Error(heading, body string) string {
  return fmt.Sprintf(`
! %v !

%v
`, strings.ToUpper(heading), wordwrap.WrapString(body, 70))
}
