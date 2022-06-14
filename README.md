## setup

- Wait for cluster to fully scale down
- `cd scenairo1`
- `./run.sh`

## narrative

Cluster will now look like this:

![2022-06-14-101145_screenshot](https://user-images.githubusercontent.com/74934191/173640105-60a0c418-3385-4f8f-9767-b5b6abecce7f.png)


Events look like this (indicating cluster is scaling up):

![2022-06-14-101215_screenshot](https://user-images.githubusercontent.com/74934191/173640198-150b55c2-d70e-4058-a470-596f69c069b3.png)

As the smalljob's finish, the cluster autoscaler kicks in an tries to schedule longjob on a different node. This is particularly strange becase longjob is actually running on a dedicated node during my testing.

![2022-06-14-101457_screenshot](https://user-images.githubusercontent.com/74934191/173640366-1a755861-941b-4f85-b527-ff1bf3423636.png)
