function __ito_root
    set -l root
    if set -q ITO_ROOT
        set root $ITO_ROOT
    else
        set root ~/.ito
    end

    if string match -qr '^~(?=/|$)' -- $root
        set root (string replace -r '^~' $HOME $root)
    end

    echo $root
end

function __ito_entries
    set -l root (__ito_root)
    if test -z "$root"
        return
    end
    if not test -d "$root"
        return
    end

    set -l entries (command find "$root" -mindepth 1 -maxdepth 1 -print 2>/dev/null | command sort)

    for entry in $entries
        set -l name (string replace -r '^.*/' '' -- $entry)
        if string match -qr '^\.' -- $name
            continue
        end
        printf '%s\n' $name
    end
end

complete -c ito -s h -l help -d 'ヘルプを表示'
complete -c ito -f -n '__fish_use_subcommand' -a list -d 'エントリを列挙'
complete -c ito -f -n 'not __fish_seen_subcommand_from list; and __fish_use_subcommand' -a '(__ito_entries)'
complete -c ito -f -n '__fish_seen_subcommand_from list'
