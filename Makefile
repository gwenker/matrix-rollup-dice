clean:
	rm -rf matrix-rollup-dice-package
	rm -f matrix-rollup-dice.zip
	rm -f matrix-rollup-dice-linux-arm-7
	
build-arm:
	env GOOS=linux GOARCH=arm GOARM=7 go build -o matrix-rollup-dice-linux-arm-7 .

package:	
	rm -rf matrix-rollup-dice-package
	mkdir -p matrix-rollup-dice-package
	cp matrix-rollup-dice-linux-arm-7 matrix-rollup-dice-package/
	cp -R fonts matrix-rollup-dice-package/
	cp -R images matrix-rollup-dice-package/
	zip -r matrix-rollup-dice.zip matrix-rollup-dice-package
