/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : stats
 * @created     : domingo ene 17, 2021 23:44:29 CST
 */

#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include "sys/types.h"
#include "sys/sysinfo.h"

typedef struct {
    double free;
    double used;
    double total;
    double swapFree;
    double swapUsed;
    double swapTotal;
} memory;

typedef struct CPUPercent{
    long long totalUser;
    long long totalUserNice;
    long long totalSys;
    long long totalIdle;
    long long totalIOWait;
    long long totalIRQ;
    long long totalSoftIRQ;
    long long prevTotalUser;
    long long prevTotalUserNice;
    long long prevotalSys;
    long long prevTotalIdle;
    long long prevTotalIOWait;
    long long prevTotalIRQ;
    long long prevTotalSoftIRQ;

    struct CPUPercent *next;
} CPUPERCENT;

typedef struct {
    long free;
    long used;
    long total;

} disk;

int ramStatus(char*, memory*);
CPUPERCENT *createNode(int);

int main(int argc, char *argv[]) {
    memory mem;
    CPUPERCENT *start = NULL, *current, *next;
    int error, i, numCPU = 0, numCPUConf = 0, starting = 0;
    char *cpu = (char*)malloc(8 * sizeof(char));
    char *useCpu = (char*)malloc(64 * sizeof(char));

    numCPU = get_nprocs();
    numCPUConf = get_nprocs_conf();

    if(numCPU == 0 || numCPUConf == 0) {
        free(cpu);
        free(useCpu);
        return EXIT_FAILURE;
    }

    // Getting status of system
    for(;;) {
        if (starting == 0) {
            for (i=-1; i<numCPUConf; i++) {
                printf("Get checkpoint data");
            }
        }

        starting++;
 
        for (i=-1; i<numCPUConf; i++) {
            /*if(!start) {
                printf("no node");
            } else {
                printf("with one node");
            }*/
            if (i == -1) {
                strcpy(cpu, "cpu");
            } else {
                sprintf(cpu, "cpu%d", i);
            }
            printf("cpu value: %s\n", cpu);
        }

        error = ramStatus("GB", &mem);

        if (error == 1) {
            perror("Error: RAM status error");
            return EXIT_FAILURE;
        }

        printf("total memory %f\n", mem.total);

        break;
    }

    return EXIT_SUCCESS;
}

// Get ram status in (MB or GB)
int ramStatus(char *units, memory *status) {
    struct sysinfo memInfo;
    double unitDivision;

    if(strcmp(units, "MB") == 0) {
        unitDivision = 1024 * 1024;
    } else if (strcmp(units, "GB") == 0) {
        unitDivision = 1024 * 1024 * 1024;
    } else {
        return EXIT_FAILURE;
    }

    // calling status function
    sysinfo(&memInfo);

    // RAM
    status->total = memInfo.totalram / unitDivision;
    status->free = memInfo.freeram / unitDivision;
    status->used = (memInfo.totalram - memInfo.freeram) / unitDivision;
    // SWAP
    status->swapTotal = memInfo.totalswap / unitDivision;
    status->swapFree = memInfo.freeswap / unitDivision;
    status->swapUsed = (memInfo.totalswap - memInfo.freeswap) / unitDivision;

    return EXIT_SUCCESS;
}

// Get swap status
// Get CPUs info
CPUPERCENT *createNode(int ncpu) {
    if(ncpu == -1) {
        printf("complete cpu");
    } else {
        printf("number of cpu");
    }
}

// Get Disk info

