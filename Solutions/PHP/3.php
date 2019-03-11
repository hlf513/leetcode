<?php

/**
 * 暴力法（子串+判重）
 *
 * 时间复杂度：n^3
 */
class Solution_One
{

	/**
	 * 通过两层循环获取所有子串
	 * - 第一层：从第一位开始，每次往后移一位
	 * - 第二层：从第一层的位置上开始，每次往后移一位
	 *
	 * @param String $s
	 *
	 * @return Integer
	 */
	public function lengthOfLongestSubstring($s)
	{
		$len = strlen($s);
		$max = 0;
		for ($i = 0; $i < $len; $i++) {
			for ($j = $i + 1; $j <= $len; $j++) {
				if ($this->uniqueString($s, $i, $j)) {
					$max = max($max, $j - $i);
				}
			}
		}

		return $max;
	}


	/**
	 * 判断子串是否有重复字符
	 *
	 * @param string $s
	 * @param int    $start
	 * @param int    $end
	 *
	 * @return bool
	 */
	private function uniqueString(string $s, int $start, int $end)
	{
		$set = [];
		for ($i = $start; $i < $end; $i++) {
			$char = $s[$i];
			if (isset($set[$char])) {
				return false;
			}
			$set[$char] = 1;
		}

		return true;
	}
}

/**
 * 滑动窗口算法
 *
 * 集合 + 双指针；判断右值
 * - 若右值在集合里
 *   - 右指针+1（右值写入集合）
 *   - 计算并更新集合中的历史最大长度
 * - 若右值不在集合里
 *   - 左指针+1（最左值从集合中删除）
 *
 * 时间复杂度：2n
 */
class Solution_Two
{
	public function lengthOfLongestSubstring(string $s)
	{
		$len = strlen($s);
		$i = $j = 0;
		$set = [];
		$max = 0;

		while ($i < $len & $j < $len) {
			if ( ! in_array($s[$j], $set)) {
				// j 往右移一位
				array_push($set, $s[$j]);
				$j++;
				$max = max($max, $j - $i);
			} else {
				// i 往右移一位
				array_shift($set);
				$i++;
			}
		}

		return $max;
	}
}

/**
 * 优化滑动窗口算法
 * 
 * 普通滑动窗口算法，i 一次移动一位，实际上 i 可以移动 j 个位置
 * 
 */
class Solution_Three
{
	public function lengthOfLongestSubstring(string $s)
	{
		$len = strlen($s);
		$max = 0;
		$hash = [];

		for ($i = 0, $j = 0; $j < $len; $j++) {
			if (isset($hash[$s[$j]])) {
				// i 跳过 j 个位置
				$i = max($hash[$s[$j]] + 1, $i);
			}
			$max = max($max, $j - $i + 1);
			// 放入 hash 表
			$hash[$s[$j]] = $j;
		}

		return $max;
	}
}

$string = "abcabcbb";
$solution = new Solution_Three();
$output = $solution->lengthOfLongestSubstring($string);

print_r($output);
