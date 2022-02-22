cur_dir=$(cd $(dirname $0);pwd)

cd $cur_dir &&
chmod +x small-tools && 
nohup ./small-tools config/conf.yml > /dev/null 2>&1 &