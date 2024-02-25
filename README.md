# JS-like slice helper functions that run in parallel and use as much hardware as possible

## TLDR

Under the hood the workload is split into batches that are evenly distributed across the CPU cores.

E. g. if your CPU has 8 cores and your array consists of 1954 items the batches would look like this:

```
[245 245 244 244 244 244 244 244] // 244 * 6 + 245 * 2 = 1954
```

## Usage

```go

input := []int{1, 2, 3, 4}

result := parallel.Map(input, func(v int) string {
	return fmt.Sprintf("-%v-", v)
})

// Output: [-1- -2- -3- -4-]
```

## Motivation

- I grew tired of implementing it for each project
- [Popular libraries](https://github.com/samber/lo/blob/7887963db823866bbe86fb8330ccb9ae5972655a/parallel/slice.go#L13-L14) just create as many workers as there are items in the input slice which is not efficient
