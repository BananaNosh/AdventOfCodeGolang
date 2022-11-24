#!/bin/bash
YEAR=$1
MAX=$2

if [ -z $YEAR ]; then YEAR=2000; fi
if [ -z $MAX ]; then MAX=25; fi

SHORT_YEAR=$(echo $YEAR | cut -c 3-4)
FOLDER="AoC$YEAR"
if ! [[ -d $FOLDER ]]; then  mkdir $FOLDER; fi

cd $FOLDER || exit
INPUT_F="input"
if ! [[ -d $INPUT_F ]]; then  mkdir $INPUT_F; fi
for i in {1..25};
do
  FILE="AoC_""$SHORT_YEAR""_$i.go"
  if ! [[ -f $FILE ]]; then
    cat ../temp | sed "s/dd/$i/" | sed "s/yyyy/$YEAR/" | sed "s/yy/$SHORT_YEAR/" > $FILE;
  fi
#  (
#  cd $INPUT_F || continue
#  FILE="input""_$i.txt"
#  if ! [[ -f $FILE ]]; then
#    touch $FILE
#  fi
#  )
done
