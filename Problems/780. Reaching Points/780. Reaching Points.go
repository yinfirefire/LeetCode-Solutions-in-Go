package main

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	//we can iterate backwards from tx and ty, the larger one in tx/ty must come from the addition of another
	//until one of them is equals to sx or sy, this is sure to happen
	for sx < tx && sy < ty {
		if tx > ty {
			tx -= ty
		} else {
			ty -= tx
		}
	}
	if tx == sx && sy <= ty && (ty-sy)%sx == 0 {
		//ty = sy+n*x
		return true
	}
	if ty == sy && sx <= tx && (tx-sx)%sy == 0 {
		//tx = sx+n*y
		return true
	}
	return false
}
