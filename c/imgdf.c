#include "imgdf.h"

int imgdf_detect(uint8_t *buf, int limit, uint8_t c){
  for(int i = 0; i < limit; i++){
    if(buf[i] == c) {
      return i;
    }
  }
  return -1;
}

int imgdf_reading(imgdf_handler_t *ih, uint8_t *buf, int limit){
  uint8_t * x = (uint8_t *)&(ih->df);
  for(int i = 0; i < limit; i++){
    x[i] = buf[i];
    ih->Num++;
  }
  return -1; // 255
}

void imgdf_hanlder_reset(imgdf_handler_t * ih){
  ih->Num = 0;
  time(&(ih->last));
}

uint8_t imgdf_get_sender(imgdf_handler_t * ih){
  if(ih->Num < 1){
    return -1; //255
  }
  return ih->df.sender;
}

uint8_t imgdf_get_receiver(imgdf_handler_t * ih){
  if(ih->Num < 2){
    return -1;//255
  }
  return ih->df.reciever;
}

uint16_t imgdf_get_dfNumber(imgdf_handler_t * ih){
  if(ih->Num < 4){
    return -1; // 65535
  }
  return ih->df.dfNumber;
}
