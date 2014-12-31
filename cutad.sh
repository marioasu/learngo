#/bin/bash

if [ ! $1 ] ; then
echo "parameters incorrect" >&2
exit 2;
fi

files_path=$1
mkdir -p $files_path/../new/
filelist=`ls $files_path`
for file in $filelist
do  gawk -F '\t' 'BEGIN{OFS="|";ORS="\n"} /[0-9]/ {print $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11}' $files_path/./$file > $files_path/../new/$file
done
echo "cut file done"
