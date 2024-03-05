# Project Introduction

Read the useragent information in the first column of the excel file, determine the device type, write the result into
the second column, and generate a new excel file.

# Quick start

Install goreleaser and execute.

    goreleaser --snapshot --clean

Enter the platform directory in the dist folder corresponding to the machine.

    cd dist/uaexcel_darwin_arm64
    ./uaexcel ../../testdata/useragent.xlsx ../../testdata/new.xlsx
