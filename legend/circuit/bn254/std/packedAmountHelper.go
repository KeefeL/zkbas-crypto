/*
 * Copyright © 2021 Zecrey Protocol
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package std

func UnpackAmount(api API, packedAmount Variable) Variable {
	amountBits := api.ToBinary(packedAmount, 40)
	mantissa := api.FromBinary(amountBits[5:]...)
	exponent := api.FromBinary(amountBits[:5]...)
	for i := 0; i < 32; i++ {
		isRemain := api.IsZero(api.IsZero(exponent))
		mantissa = api.Select(isRemain, api.Mul(mantissa, 10), mantissa)
		exponent = api.Select(isRemain, api.Sub(exponent, 1), exponent)
	}
	return mantissa
}

func UnpackFee(api API, packedFee Variable) Variable {
	amountBits := api.ToBinary(packedFee, 16)
	mantissa := api.FromBinary(amountBits[5:]...)
	exponent := api.FromBinary(amountBits[:5]...)
	for i := 0; i < 32; i++ {
		isRemain := api.IsZero(api.IsZero(exponent))
		mantissa = api.Select(isRemain, api.Mul(mantissa, 10), mantissa)
		exponent = api.Select(isRemain, api.Sub(exponent, 1), exponent)
	}
	return mantissa
}
