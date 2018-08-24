#!/bin/bash
counter=0

FILE=hello.go
OLDTIME=2


echo "Powered by: https://gobyexample.com/ "
echo "Press [CTRL+C] to stop.."
sleep 2
clear

while true
do    
    CURTIME=$(date +%s)
    FILETIME=$(stat $FILE -c %Y)
    TIMEDIFF=$(expr $CURTIME - $FILETIME)

    if [ $TIMEDIFF -lt $OLDTIME ]; then
        ((counter++))
	    echo "Running \"hello.sh\" & Building" $FILE: Livebuild \#$counter
        go build
        echo "-------------Running #"$counter"--------------"
        ./hello
        echo "-------------Ended  #"$counter"---------------"
        sleep 15
    else
        clear
        echo "Nothing to do.."
        sleep 2
        
    fi
done
