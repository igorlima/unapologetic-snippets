" vim -N -u array.vimrc
" http://localhost:4000/docs/languages/vim/script-for-debugging
"
" ------
" BASICS
" ------
let myArray=['one', 'two', 'three']

echom myArray[0]
echom "is `one`"

echom myArray[-1]
echom "is `three`"

" ---------
" FUNCTIONS
" ---------
call add(myArray, 'four')
echom "adds a new elements."
let myArray+=['five']
echom "also works"

echom get(myArray, 0, 'default value')
echom "reads a value, with a fallback"

echom len(myArray)
echom "returns 5"

echom index(myArray, 'three')
echom "returns `2` (or `-1` if not found)"

echom join(myArray, '/')
echom "returns `one/two/three/four/five`"

echom split('one/two/three', '/')
echom "creates `['one', 'two', 'three']`"

" ---------
" REFERENCE
" ---------
echom " "
echom "Array (List) functions in vimscript"
echom " https://blog.pixelastic.com/2023/02/07/array-functions-in-vimscript/"
echom " "
echom "Learn Vimscript the Hard Way"
echom " https://learnvimscriptthehardway.stevelosh.com/"
echom " "
echom "Lists"
echom " https://learnvimscriptthehardway.stevelosh.com/chapters/35.html"
echom " "
echom "String Functions"
echom " https://learnvimscriptthehardway.stevelosh.com/chapters/27.html"

