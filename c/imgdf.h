#ifndef github_com_kranfix_imgdf_c
#define github_com_kranfix_imgdf_c

#include <stdint.h>
#include <time.h>

typedef struct  {
  uint8_t sender;
  uint8_t reciever;
  uint16_t dfNumber;
  uint8_t buf[60];
} imgdf_t;

typedef struct {
  int Num;
  time_t last;
  imgdf_t df;
} imgdf_handler_t;

/*
 Return the first place where is a character @c in
 a  buffer @buf with a @limit
*/
int imgdf_detect(uint8_t *buf, int limit, uint8_t c);

/*

*/
void imgdf_hanlder_reset(imgdf_handler_t * ih);
uint8_t imgdf_get_sender(imgdf_handler_t * ih);
uint8_t imgdf_get_receiver(imgdf_handler_t * ih);
uint16_t imgdf_get_dfNumber(imgdf_handler_t * ih);

#endif
