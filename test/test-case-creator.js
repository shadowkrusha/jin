const fs = require('fs');

var exec = require('child_process').execFile;

let jsonFile  = 'test/test-json.json';
let pathFile  = 'test/test-json-paths.json';
let valueFile = 'test/test-json-values.json';

// clearing files
fs.writeFileSync(pathFile, "");
fs.writeFileSync(valueFile, "");

// console.log(process.argv[2])

var mainArray = JSON.parse(fs.readFileSync(jsonFile));
// console.log(mainArray)

const pathToString = (arr) => {
	var result = "["
	arr.forEach((e) => {
		result += JSON.stringify(e) + ","
	})
	result = result.slice(0, result.length - 1)
	result += "]"
	return result
}

function createPaths(obj, key) {
	const keys = Object.keys(obj)
	const values = Object.values(obj)
	for ( var i = 0 ; i < keys.length ; i ++ ){
		key.push(keys[i])
		let path = pathToString(key)
		fs.appendFileSync(pathFile, path + `\n`)
		key.pop()
		if (typeof values[i] === 'object' && values[i] != null && values[i] != undefined) {
			key.push(keys[i])
			createPaths(values[i], key)
			key.pop()
		}
	}
	
}

function createValues(obj) {
	const keys = Object.keys(obj)
	const values = Object.values(obj)
	for ( var i = 0 ; i < keys.length ; i ++ ){
		fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`)
		if (typeof values[i] === 'object' && values[i] != null && values[i] != undefined) {
			createValues(values[i])
		}
	}
}


function setValues(obj, val) {
	const keys = Object.keys(obj)
	const values = Object.values(obj)
	for ( var i = 0 ; i < keys.length ; i ++ ){
		if (typeof values[i] === 'object') {
			setValues(values[i], val)
		}else{
			obj[keys[i]] = val
		}
	}
	return obj
}

createPaths(mainArray, [])
createValues(mainArray, [])