//217. Contains Duplicate
func containsDuplicate(nums []int) bool {
	hash := make(map[int]int);
  
	for _, val := range nums{
	  if _, isFound := hash[val]; isFound {
		  return true;
	  }
	  hash[val] = 1;
	}
	  return false;
  }