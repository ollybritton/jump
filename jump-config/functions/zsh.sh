jump () {

    # jump takes the first command-line argument as an alias, searches jump-config for a directory associated with it, and changes the directory accordingly.
    # INFO: colors=yes, shell=zsh

    # Find the initial result of tring to parse an alias.
    local jump_result=$(jump-config parse $1)

    # jump_result may be a path, "MULTIPLE_MATCH" (saying there's more than one directory matching that alias) or "NO_MATCH" (saying that no alias matching that pattern could be found)

    if [[ $jump_result == "NO_MATCH" ]]
    then
    echo "There are no directories associated with that alias."

    elif [[ $jump_result == "NO_ARGS" ]]
    then
    echo "You need to give an alias to navigate to, like this"
    echo "jump [alias]"

    elif [[ $jump_result == "MULTIPLE_MATCH" ]]
    then

    echo "More than one directory matches that alias."

    # This command will return a list of matches in the following format:
    # [alias1]|[directory1]++[alias2]|[directory2] ...
    matches=$(jump-config list $1)

    # The following code reads the matches into an array.
    IFS='>'
    read -rA ALL_MATCHES_LIST <<< "$matches"

    # Loop through all matches, split them up and print them out in a human-readable way.
    for ((i = 1; i <= $#ALL_MATCHES_LIST; i++)); do
        value=${ALL_MATCHES_LIST[i]}

        IFS='|'
        read -rA MATCH <<< "$value"

        echo "($i) ${MATCH[1]} -> ${MATCH[2]}"
    done

    # Ask the user for an index of which directory they wish to navigate to.
    printf "\nPlease enter directory you wish to go to: "
    read directoryIndex
    chosen=${ALL_MATCHES_LIST[$directoryIndex]}

    # Split the chosen option up, and navigate to the path given.
    IFS='|'
    read -rA DESIRED_MATCH <<< "$chosen"
    eval cd ${DESIRED_MATCH[2]}

    else
    eval cd $jump_result
fi
}