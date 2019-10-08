dirbase <- "/kiwy/resources/csv"
files <- unlist(list.files(dirbase))

for(i in 1:length(files)) {
  data <- read.csv(paste(dirbase, files[i], sep="/"), header=FALSE)
  names(data) <- c('table', 'key', 'startDate', 'endDate', 'sysDate', 'sysTime')

  sysTimes <- unlist(data['sysTime'])
  
  percentiles <- quantile(sysTimes, prob = c(0.50, 0.75, 0.90, 0.95, 0.99))
  filename <- paste("/kiwy", "resources", "percentiles", files[i], sep="/")

  write.csv(percentiles, file = filename)
}
