clean:
	rm -rf __debug_bin \;
	rm -rf gin-bin \;

swagger:
	swag init --generalInfo server/routes/swagger.go --parseDependency --parseInternal

run:
	sh ./scripts/run.sh

rename:
	sh ./scripts/rename.sh