package vmgroupcalc

import (
    "github.com/kun-lun/digester/pkg/common"
)

type Requirment struct {
    ConcurrentUserNumber int
}

func Calc(r Requirment) common.IaaS {
    res := common.IaaS{
        Size: common.SizeSmall,
    }
    x := r.ConcurrentUserNumber
    if x >= 1000 {
        res.Size = common.SizeMedium
    }
    if x >= 2000 {
        res.Size = common.SizeLarge
    }
    if x >= 4000 {
        res.Size = common.SizeMaximum
    }

    return res
}
