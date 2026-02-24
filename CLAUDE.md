# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

**fortune** generates a daily email containing the joke-of-the-day, fire department shift information, and elapsed time data. It creates an email file in MailEnable's queue format.

**Module**: `github.com/pootwaddle/fortune`
**Type**: Executable
**Dependencies**: `dadjoke`, `dayplus`, `ljemail`, `shift`, `slogger`

## Processing

1. Load jokes from `c:/autojob/fortune.dat`
2. Select joke-of-the-day based on current date
3. Determine current FD shift via `shift.GetShift()`
4. Calculate elapsed time data via `dayplus.ElapsedTime()`
5. Build email with `ljemail` (headers, body, footer)
6. Write `.MAI` file to MailEnable pickup queue

## Scheduled Execution

Runs daily at 03:33:35 via `schedule.exe` -> `FORTUN.BAT`.

## Build Commands

```powershell
go build
```
