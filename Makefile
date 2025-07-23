all: build

.PHONY:fmt build release github-release clean

# format golang code
fmt:
	@if command -v goimports >/dev/null 2>&1; then \
		find ./app -name "*.go" -exec goimports -local github.com/wiselike/leanote-of-unofficial -l -w {} \; ;\
	else \
		find ./app -name "*.go" -exec go fmt {} \; ;\
	fi
	@simple-formater -dir app
	@simple-formater -dir public

# only build temporarily
build:
	@rm -rf tmp/
	CGO_ENABLED=0 revel build . tmp/

# build js
gulp:
	@cd public/tinymce; rm -f tinymce.js tinymce.dev.js tinymce.min.js tinymce.jquery.dev.js tinymce.full.js tinymce.full.min.js
	@cd public/tinymce; grunt minify;
	@cd public/tinymce; grunt bundle --themes=leanote --plugins=autolink,link,leaui_image,leaui_mindmap,lists,hr,paste,searchreplace,leanote_nav,leanote_code,tabfocus,table,directionality,textcolor;
	@gulp

# build all and rerun leanote
release: gulp
	@rm -rf release/
	CGO_ENABLED=0 revel build . release/
	rsync -azr --delete --delete-before --exclude github.com/wiselike/leanote-of-unofficial/conf/app.conf --exclude github.com/wiselike/leanote-of-unofficial/public/upload --exclude github.com/wiselike/leanote-of-unofficial/mongodb_backup -e 'ssh -p 22' release/src/ root@192.168.0.12:/root/dockers/leanote/leanote/src
	rsync -azr release/leanote-of-unofficial  -e 'ssh -p 22' root@192.168.0.12:/root/dockers/leanote/leanote/leanote-of-unofficial
	rm -rf release/
	ssh -p 22 root@192.168.0.12 "docker restart leanote"

github-release: gulp
	@rm -rf release github-release;
	@mkdir github-release;
	CGO_ENABLED=0 GOARCH=arm64 revel build . release/
	@mv release/leanote-of-unofficial github-release/linux-arm64-leanote-of-unofficial
	tar czf js-release.tar.gz release && tar cJf js-release.tar.xz release;
	@mv js-release.tar.gz github-release && mv js-release.tar.xz github-release;
	CGO_ENABLED=0 GOARCH=amd64 revel build . release/
	@mv release/leanote-of-unofficial github-release/linux-x64-leanote-of-unofficial
	@rm -rf release;
	@echo -e "\n\ngithub-release finished in ./github-release:" && ls -alh github-release;

clean:
	rm -rf tmp/ release/ github-release target
