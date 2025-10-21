package errorsx

import "errors"

var ErrWrongNumberOfElements = errors.New("недостаточно элементов")
var ErrNegativeSteps = errors.New("количество шагов меньше или равно нулю")
var ErrNegativeDuration = errors.New("отрицательное или нулевое время")
var ErrNegativeWeight = errors.New("вес меньше или равен нулю")
var ErrNegativeHeight = errors.New("высота меньше или равна нулю")
var ErrUnknownAcrivityType = errors.New("неизвестный тип тренировки")
