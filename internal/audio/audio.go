// Copyright 2015 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !js

package audio

const SampleRate = 44100

func Init() {
	// TODO: Implement
}

func Start() {
	// TODO: Implement
}

func Append(channel int, l []float32, r []float32) bool {
	// TODO: Implement
	return false
}

func CurrentBytes() int {
	// TODO: Implement
	return 0
}

func Update() {
	// TODO: Implement
}