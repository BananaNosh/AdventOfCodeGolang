#!/bin/bash
YEAR=$1
MAX=$2


if [ -z $YEAR ]; then YEAR=2000; fi
if [ -z $MAX ]; then MAX=25; fi

SHORT_YEAR=$(echo $YEAR | cut -c 3-4)
FOLDER="AoC$YEAR"
if ! [[ -d $FOLDER ]]; then  mkdir $FOLDER; fi

# uncomment lines in main file
SWITCH_LINE_OFFSET=22
LINE_START=$(echo "$SWITCH_LINE_OFFSET + 2*($MAX+1)" | bc)
BEFORE_LINE_START=$(echo "$LINE_START - 1" | bc)
LINE_END=$(echo "$SWITCH_LINE_OFFSET+51" | bc)
echo $MAX $LINE_START $LINE_END
if [ $MAX -lt 25 ]; then
  sed -i "$LINE_START,$LINE_END s/^\t\([^\/\/]\)/\t\/\/\1/" AoC_main.go
fi
if [ $MAX -gt 1 ]; then
  sed -i -r "$SWITCH_LINE_OFFSET,$BEFORE_LINE_START s/^\t\/\//\t/" AoC_main.go
fi

cd $FOLDER || exit
INPUT_F="input"
if ! [[ -d $INPUT_F ]]; then  mkdir $INPUT_F; fi
for i in $(seq 1 $MAX);
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
