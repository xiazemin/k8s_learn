ls |grep yaml |grep -v redis |xargs -I{} kubectl apply -f {}