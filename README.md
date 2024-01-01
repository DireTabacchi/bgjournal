# Blood Glucose Journal

`bgjournal` is a simple journaling app to keep track of a diabetic's
blood-glucose level, as well as provide some basic statistics about such
levels, including average daily levels and insulin use.

## Installation

Ensure that you have the `go` tool installed before trying to install
`bgjournal`. If you need to install the `go` tool, visit the
[Go website](go.dev) and follow the instructions to download and install the
`go` tool.

Since this app is in development, cloning the repository to your computer with

```
git clone https://github.com/DireTabacchi/bgjournal.git
```

Changing directory (`cd bgjournal`) and running the Go compiler

```
go build .
```

will give you the current working version of the application.

Currently, the application only contains tests for the implemented features.

## Milestones/Features

- [X] File writing.
- [X] File reading.
- [ ] User can interactively create new entries.
- [ ] User can interactively query entries.
- [ ] Statistics are displayed for each entry.
    - [ ] Current daily average.
    - [ ] Current weekly average (for a day's week).
    - [ ] Weekly average compared to previous week.
    - [ ] Weekly BG checks per day.
- [ ] Statistics displayed for a selected month.
- [ ] Statistics compared to the previous month.

## Feature Requests/Bugs

If you'd like to request a feature not shown in the **Milestones/Features**
section, please raise an issue on this repository. If you have found a bug,
please raise an issue on this repository, with the following information:

- Date of bug
- Steps to recreate bug (What did you do when the bug occurred?)
- Error given at encounter

As this application is somewhat of a personal/passion project, please do not
expect all features/bugs to be addressed immediately. Major bugs will be
adressed most immediately, and agreeable features will be considered.
