#!/usr/bin/python                                        
                                                         
import sys
                                                         
MAX_MORSE_CHAR_LEN = 6                                                                                             
MAX_CONTEXT_WORD_LEN = 10                                
MAX_MORSE_WORD_LEN = 80
                                                         
class morseTable:                                        
                                                         
    def __init__(self):                                  
        self.alpha_to_morse = dict()
        self.context_words = list()                      
        self.morse_words = list()                        
                                                         
        num_lines = sys.stdin.readline().strip(" \n")
        self._parse_table_lines()                        
        self._parse_context_line()                       
        self._parse_morse_line()                         
                                                         
    def _get_line(self):
        while True:                                      
            line = sys.stdin.readline().strip(" \n")     
            if not line: break                           
            if len(line) < 1: continue                   
            if line == '*': break
                                                                                                                   
            return line                                  
        return False
                                                         
    def _parse_table_lines(self):                        
        while True:                                      
            current_line = self._get_line()              
            if not current_line: break
                                                         
            letter, morse = current_line.split()         
            self.alpha_to_morse[letter] = morse[:MAX_MORSE_CHAR_LEN]
        return                                           

    def _parse_context_line(self):                       
        while True:                                      
            current_line = self._get_line()              
            if not current_line: break                   

            if len(current_line) <= MAX_CONTEXT_WORD_LEN:                                                          
                self.context_words.append(current_line)  
        return                                           

    def _parse_morse_line(self):                         
        while True:                                      
            current_line = self._get_line()              
            if not current_line: break                   

            words_to_decode = [word                      
                for word in current_line.split()         
                if len(word) <= MAX_MORSE_WORD_LEN       
            ]                                            

            self.morse_words.extend(words_to_decode)                                                      
        return                                                                                    

if __name__ == "__main__":                               
    currentTable = morseTable()
