'use strict';

const readline = require('readline');
let line_reader = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
    terminal: false
});

let num_lines = null;
let letter_to_morse = {};
let context_words = [];
let morse_words = [];

const MAX_MORSE_CHAR_LEN = 6;
const MAX_CONTEXT_WORD_LEN = 10;
const MAX_MORSE_WORD_LEN = 80;

let parse_phase = 0;
const PARSE_PHASES = {
    ENCODE: 0,
    CONTEXT: 1,
    DECODE: 2
};

line_reader.on('line', line => {
    let trimmed_line = line.trim();
    // First line of the input indicates the number of lines that follow
    if(num_lines == null) {
        num_lines = parseInt(trimmed_line, 10);
        return;
    }
    // Asterisk characters indicate the end of the current phase
    if( trimmed_line === '*' ) {
        parse_phase++;
        return;
    }
    if(parse_phase === PARSE_PHASES.ENCODE) {
        return parse_table_line(trimmed_line);
    } else if (parse_phase === PARSE_PHASES.CONTEXT) {
        return parse_context_line(trimmed_line);
    } else if (parse_phase === PARSE_PHASES.DECODE) {
        return parse_morse_line(trimmed_line);
    } else {
        return;
    }
});

line_reader.on('close', () => {
    decode();
});

function parse_table_line(line) {
    if(!line.length) {
        return;
    }
    // an uppercase letter or a digit C, zero or more blanks,
    // and a sequence of no more than six periods and hyphens
    let letter = line[0];
    let morse = line.substring(1).trim().substring(0, MAX_MORSE_CHAR_LEN);
    letter_to_morse[letter] = morse;
}

function parse_context_line(line) {
    //one word per line, possibly preceded and followed by blanks
    if(line.length <= MAX_CONTEXT_WORD_LEN) {
        context_words.push(line);
    }
}

function parse_morse_line(line) {
    let words_to_decode = line.split(/\s+/)
        .filter(morse_word => morse_word.length && morse_word.length <= MAX_MORSE_WORD_LEN);
    //morse words separated by blanks or end-of-line characters
    morse_words = morse_words.concat(words_to_decode);
}

function decode() {
    // write your code here and
    // console.log the expected output
}
