```
mkdir vgop
cd vgop/
go mod init vgop
git init
brew upgrade kubebuilder


kubebuilder init --domain kubotal.io
kubebuilder edit --multigroup=true

kubebuilder create api --group vgop --version v1alpha1 --kind LVMVolumeGroup
Create Resource [y/n]
y
Create Controller [y/n]
y

make manifests

git commit -m "Initial commit"
```
