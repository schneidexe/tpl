if [ ! -d "${DST_PATH}" ]; then
  mkdir -p ${DST_PATH}
fi
for file in $(find ${SRC_PATH} -maxdepth ${DEPTH} -type f -printf '%P\n'); do
  tpl -t ${SRC_PATH}/${file} > ${DST_PATH}/${file}
done
