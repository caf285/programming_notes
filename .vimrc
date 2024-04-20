" for macOS
syntax on

" golang color
so ~/.vim/syntax/go.vim

" turn on mouse
set mouse=a
map <ScrollWheelUp> 3<Up>
map <ScrollWheelDown> 3<Down>
imap <ScrollWheelUp> <Up>
imap <ScrollWheelDown> <Down>

" turn on lines
set number

" quick i/o for mouse and lines
nnoremap <C-M><C-I> :set mouse=a<CR>:set number<CR>
nnoremap <C-M><C-O> :set mouse=c<CR>:set nonumber<CR>

" set wrap for <,>,h,l,[,] keys
set whichwrap+=<,>,h,l,[,]

" fix tab
set tabstop=2
set softtabstop=2
set expandtab

" set cursorline and cursorcolumns
set cursorline

" fix auto insertion of comment blocks
set formatoptions-=cro

" fix colors
:colorscheme slate
hi Normal ctermfg=white ctermbg=NONE
hi Visual cterm=bold ctermfg=white ctermbg=red
hi NonText cterm=bold ctermfg=darkred ctermbg=NONE
hi Comment ctermfg=darkgreen ctermbg=NONE
hi Constant ctermfg=darkgray ctermbg=NONE
hi CursorLineNr cterm=bold ctermfg=darkred ctermbg=NONE
hi LineNr ctermfg=darkgray ctermbg=NONE
hi Error ctermfg=red ctermbg=NONE
hi Statement cterm=bold ctermfg=darkred ctermbg=NONE
hi Special ctermfg=cyan ctermbg=NONE
hi PreProc ctermfg=yellow ctermbg=NONE
hi Identifier ctermfg=darkcyan ctermbg=NONE
hi Search ctermfg=black ctermbg=cyan
hi MatchParen cterm=bold ctermfg=white ctermbg=red
hi Todo cterm=bold ctermfg=white ctermbg=red
hi CursorLine cterm=bold ctermbg=black ctermfg=NONE
hi CursorColumn cterm=bold ctermbg=black ctermfg=NONE

" =========================( Vundle )
" -----( setup)
" in bash, run `git clone https://github.com/VundleVim/Vundle.vim.git ~/.vim/bundle/Vundle.vim`
" in .vimrc (this file), run `:PluginInstall` 
" /----( END setup)

set nocompatible              " be iMproved, required
filetype off                  " required

" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

" plugins
Plugin 'preservim/nerdtree'

" All of your Plugins must be added before the following line
call vundle#end()            " required
filetype plugin indent on    " required

" toggle plugins
nnoremap <C-n> :NERDTreeToggle<CR>
