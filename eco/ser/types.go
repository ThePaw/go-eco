// Copyright 2012 The Eco Authors. All rights reserved. See the LICENSE file.

package ser

type ObjFn func(sim Matrix64, p IntVector) float64
type Sampler func(mtx Matrix64, sEffort float64) IntMatrix
type OptMethod func(Matrix64, ObjFn, bool) (float64, IntVector)
