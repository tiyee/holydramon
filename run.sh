#!/usr/bin/env bash
workspace=$(cd $(dirname $0) && pwd -P)
app=holydramon
killall  $app
sleep 3


helpFunction()
{
   echo ""
   echo "Usage: $0 -c config -d running as a daemon"
   echo -e "\t-d Description of what is parameterB"
   echo -e "\t-c Config file path thar be relative to the application root directory."
   exit 1 # Exit script after printing help
}
while getopts "c:d" opt
do
   case "$opt" in
      c ) parameterC="$OPTARG" ;;
      d ) parameterD='daemon' ;;
      ? ) helpFunction ;; # Print helpFunction in case parameter is non-existent
   esac
done

# Print helpFunction in case parameters are empty
if  [ -z "$parameterC" ]
then
   echo "config parameter is empty";
   helpFunction
fi
echo "${workspace}/${parameterC}"
if  [ ! -z "$parameterD" ] && [ $parameterD=='daemon' ]
then
  nohup ${workspace}/bin/${app} -c "${workspace}/${parameterC}" >> ${workspace}/logs/output.log 2>&1 &
else
  echo "running..."
  ${workspace}/bin/${app} -c "${workspace}/${parameterC}"
fi

# Begin script in case all parameters are correct
#echo "$parameterD"
#echo "$parameterC"