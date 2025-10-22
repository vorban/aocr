# AOCR

A simple tool to generate completion badges for the Advent of Code.

## Preview

<!-- AOCR_BADGES_START -->
<img src="https://img.shields.io/badge/total_stars%20⭐-101%2F548-fcd34d?style=for-the-badge">
<br>
<img src="https://img.shields.io/badge/2015%20⭐-02%2F50-e5e5e5">
<img src="https://img.shields.io/badge/2016%20⭐-50%2F50-fcd34d">
<img src="https://img.shields.io/badge/2017%20⭐-02%2F50-e5e5e5">
<img src="https://img.shields.io/badge/2018%20⭐-00%2F50-a8a29e">
<img src="https://img.shields.io/badge/2019%20⭐-00%2F50-a8a29e">
<br>
<img src="https://img.shields.io/badge/2020%20⭐-00%2F50-a8a29e">
<img src="https://img.shields.io/badge/2021%20⭐-00%2F50-a8a29e">
<img src="https://img.shields.io/badge/2022%20⭐-00%2F50-a8a29e">
<img src="https://img.shields.io/badge/2023%20⭐-00%2F50-a8a29e">
<img src="https://img.shields.io/badge/2024%20⭐-00%2F50-a8a29e">
<br>
<img src="https://img.shields.io/badge/2025%20⭐-24%2F24-fcd34d">
<img src="https://img.shields.io/badge/2026%20⭐-23%2F24-e5e5e5">
<!-- AOCR_BADGES_END -->

## Details

The Advent of Code started in December 2015. At the time,
25 puzzles were made available between December 1st and December 25th,
allowing participants to obtain 50 stars.

In 2025 it was announced there will be only 12 puzzles in December.
This repository therefore assume 24 stars obtainable from 2025 onward.

The badge for the current year is only generated if the program
is executed during December.

This program places html markers inside your readme to overwrite
its output on successive executions.

There are 3 scenarios:

1. no readme file found in the working directory -> creates an empty one
2. readme file found without markers -> prepends badges to the file
3. readme file found with markers -> overwrites the badges

## Todolist that you know won't be done for the next 3 years

- [ ] `stdout` flag to dump the badges to the stdout
- [ ] better readme existence check - be case insensitive
- [ ] serialize stars obtained in the produced html
- [ ] update mode - change only a specific year
- [ ] show previous values in the prompt

## FAQs

### Why not automate day initialization with code templates ?

It seems too much of a time investment to handle all possible
file structures depending on preference and languages.

### Why not automate input fetching ?

Because going to the website every day and looking at the input
and examples is a little ritual I am quite fond of.

## License

BSD 3-Clause License, see [LICENSE](./LICENSE).

## Advent of Code Automation Policy Compliance

This tool does not make network requests.
