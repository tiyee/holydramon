#!/usr/bin/env bash
helpFunction()
{
   echo ""
   echo "Usage: $0 -c config -d running as a daemon"
   echo -e "\t-d Description of what is parameterB"
   echo -e "\t-c Config file path thar be relative to the application root directory."
   exit 1 # Exit script after printing help
}
while getopts "d:c:" opt
do
   case "$opt" in
      d ) parameterD='daemon' ;;
      c ) parameterC="$OPTARG" ;;
      ? ) helpFunction ;; # Print helpFunction in case parameter is non-existent
   esac
done

# Print helpFunction in case parameters are empty
if  [ -z "$parameterC" ]
then
   echo "config parameter is empty";
   helpFunction
fi

if  [ $parameterD=='daemon' ]
then
  echo "111111"
else
  echo "hacker"
fi

# Begin script in case all parameters are correct
echo "$parameterD"
echo "$parameterC"