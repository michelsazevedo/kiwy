dirbase <- "/kiwy/resources/csv"
files <- unlist(list.files(dirbase))

filename <- paste("/kiwy", "resources", "percentiles", "percentiles.csv", sep="/")

for(i in 1:length(files)) {
  data <- read.csv(paste(dirbase, files[i], sep="/"), header=FALSE)
  names(data) <- c('datetime', 'time', 'count')

  sysTimes <- round(unlist(data['time'])*1000)
  
  quantiles <- quantile(sysTimes, prob = c(0.50, 0.75, 0.90, 0.95, 0.99))

  throughput <- strsplit(files[i], ".csv")[[1]]

  throughputs <- c(replicate(5, throughput))
  percentiles <- c("p50", "p75", "p90", "p95", "p99")

  data <- data.frame(throughputs, percentiles, quantiles)

  write.table(data, filename, append = T, col.names=FALSE)
}
