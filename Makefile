# There isn't a way to automatically remove outdated files at this time
# https://github.com/OpenAPITools/openapi-generator/issues/15669
# Diffing the tracked FILES with `comm` allows us to work around this
%-client:
	test -d $*
	cp $*/.openapi-generator/FILES $*/FILES.tmp || touch $*/FILES.tmp
	make -C $* client
	@echo "Removing outdated files:" && cd $* && comm -23 FILES.tmp .openapi-generator/FILES
	cd $* && comm -23 FILES.tmp .openapi-generator/FILES | xargs -n10 rm
	rm $*/FILES.tmp
