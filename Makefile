yamlfmt.1: yamlfmt.1.md
	pandoc yamlfmt.1.md -s -t man > yamlfmt.1  

.PHONY: debian
debian:
	nfpm -f .nfpm.yaml pkg -t yamlfmt.deb
