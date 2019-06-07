  if executable('go-langserver-gtags')
    au User lsp_setup call lsp#register_server({
        \ 'name': 'go-langserver-gtags',
        \ 'cmd': {server_info->['go-langserver-gtags']},
        \ 'whitelist': ['c', 'h', 'cpp', 'cxx', 'cc', 'objc', 'objcpp'],
        \ })
  endif
