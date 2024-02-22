##########Formatting code#########################################
go fmt ./ && go vet ./
##########Compiling gva#######################################
go build -o gva cmd/main.go
##########Successfully compiled gva###########################
##########Initializing data using gva#########################
chmod +x gva
gva initdb
##########Use gva to initialize data successfully#############
##########Deleting gva########################################
rm -f gva
##########Deleting gva successfully###########################