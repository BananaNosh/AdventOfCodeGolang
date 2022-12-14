#!/bin/bash
YEAR=$1
MAX=$2


if [ -z $YEAR ]; then YEAR=2000; fi
if [ -z $MAX ]; then MAX=25; fi

SHORT_YEAR=$(echo $YEAR | cut -c 3-4)
FOLDER="AoC$YEAR"
if ! [[ -d $FOLDER ]]; then  mkdir $FOLDER; fi

# uncomment lines in main file
SWITCH_LINE_OFFSET=47
LINE_START=$(echo "$SWITCH_LINE_OFFSET + 2*($MAX+1)" | bc)
BEFORE_LINE_START=$(echo "$LINE_START - 1" | bc)
LINE_END=$(echo "$SWITCH_LINE_OFFSET+51" | bc)

IMPORT_LINE_OFFSET=3
IMPORT_LINE_START=$(echo "$IMPORT_LINE_OFFSET + ($MAX+1)" | bc)
IMPORT_BEFORE_LINE_START=$(echo "$IMPORT_LINE_START - 1" | bc)
IMPORT_LINE_END=$(echo "$IMPORT_LINE_OFFSET+25" | bc)

echo $MAX
echo $LINE_START $LINE_END $BEFORE_LINE_START
echo $IMPORT_LINE_START $IMPORT_LINE_END $IMPORT_BEFORE_LINE_START
if [ $MAX -lt 25 ]; then
  echo
  sed -i "$LINE_START,$LINE_END s/^\t\t\([^\/\/]\)/\t\t\/\/\1/" AoC_main.go
  sed -i "$IMPORT_LINE_START,$IMPORT_LINE_END s/^\t\([^\/\/]\)/\t\/\/\1/" AoC_main.go
fi
if [ $MAX -gt 1 ]; then
  sed -i -r "$SWITCH_LINE_OFFSET,$BEFORE_LINE_START s/^\t\t\/\//\t\t/" AoC_main.go
  sed -i -r "$IMPORT_LINE_OFFSET,$IMPORT_BEFORE_LINE_START s/^\t\/\//\t/" AoC_main.go
fi

cd $FOLDER || exit
INPUT_F="input"
if ! [[ -d $INPUT_F ]]; then  mkdir $INPUT_F; fi
EXAMPLE_F="example"
if ! [[ -d $EXAMPLE_F ]]; then  mkdir $EXAMPLE_F; fi
for i in $(seq 1 $MAX);
do
  SUBFOLDER="AoC_""$SHORT_YEAR""_$i"
  if ! [[ -d $SUBFOLDER ]]; then  mkdir $SUBFOLDER; fi
  cd $SUBFOLDER || exit
  FILE="AoC_""$SHORT_YEAR""_$i.go"
  if ! [[ -f $FILE ]]; then
    cat ../../temp | sed "s/dd/$i/" | sed "s/yyyy/$YEAR/" | sed "s/yy/$SHORT_YEAR/" > $FILE;
  fi
  cd ..
#  (
#  cd $INPUT_F || continue
#  FILE="input""_$i.txt"
#  if ! [[ -f $FILE ]]; then
#    touch $FILE
#  fi
#  )
done

# TODO make every day own package
