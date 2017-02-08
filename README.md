# filemerger
A filemerger is someone who will merge duplicated files on your disk. Basically it will
1. calculate every file's md5 in your disk(not whole disk but specific directory recommanded)
2. store file name(absolute path) and md5 in a key-value (filename-md5) store
3. store md5 and count in a key-value(md5-count) store
4. if a md5 shows twice or more, get the filenames and make result
5. show results and let user decide which file will be removed
