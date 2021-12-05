我们这里的Kind是CronJob了，要注意的是.spec.schedule字段是必须填写的，用来指定任务运行的周期，格式就和crontab一样，另外一个字段是.spec.jobTemplate, 用来指定需要运行的任务，格式当然和Job是一致的。还有一些值得我们关注的字段.spec.successfulJobsHistoryLimit和.spec.failedJobsHistoryLimit，表示历史限制，是可选的字段。它们指定了可以保留多少完成和失败的Job，默认没有限制，所有成功和失败的Job都会被保留。然而，当运行一个Cron Job时，Job可以很快就堆积很多，所以一般推荐设置这两个字段的值。如果设置限制的值为 0，那么相关类型的Job完成后将不会被保留。

% kubectl apply -f deployment/CronJob.yaml
error: unable to recognize "deployment/CronJob.yaml": no matches for kind "CronJob" in version "batch/v2alpha1"

 batch/v2alpha1
 apiVersion: apps/v1

 https://stackoverflow.com/questions/67520866/no-matches-for-kind-cronjob-in-version-batch-v1

 batch/v1

 % kubectl apply -f deployment/CronJob.yaml
cronjob.batch/alpine-cronjob created


% kubectl get cronjob 
NAME             SCHEDULE      SUSPEND   ACTIVE   LAST SCHEDULE   AGE
alpine-cronjob   */1 * * * *   False     0        26s             62s


% kubectl get pod     
NAME                                        READY   STATUS      RESTARTS   AGE
alpine-cronjob-27311810-f9qbb               0/1     Completed   0          10s


https://www.cnblogs.com/zjz20/p/14063389.html