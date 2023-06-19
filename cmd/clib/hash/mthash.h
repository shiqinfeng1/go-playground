/***
 * @FilePath     : mthash.h
 * @Description  : Hash calculation FPGA driver API file. See the description document for details.
 *                 The same server supports up to 5 hash board devices.
 * @Author       : lishen
 * @Version      : V2.1.1
 * @Date         : 2022-06-10 09:30:44
 * @LastEditTime : 2022-10-28 10:30:16
 * @email        : listen@jmail.com
 * @
 * @CopyRight    : Copyright (c) 2022 by MatrixTime Digit Technology Co. Ltd, All Rights Reserved.
 */
#ifndef __MTHASH_H
#define __MTHASH_H

#ifdef __cplusplus
extern "C"
{
#endif

#include <stdint.h>

    /***
     * @description: Input or output data structure
     * @param :uint32_t len: data length
     * @param :uint32_t *data: data address
     */
    typedef struct _Mt_Data_Type
    {
        uint32_t len;
        uint8_t *data;
    } Mt_Data_Type_s;

    /***
     * @description: Input parameter data structure
     * @param :uint32_t hashLen: input hash key length, should be 16
     * @param :uint8_t *hashKey: input hash key buffer, 16 bytes size
     * @param :uint32_t polyLen: input Irreducible polynomial key length, should be 16 or 0.
     *                        If it is 0, the default parameter will be used for hash calculation;
     *                        if it is 16, the input parameter will be used for hash calculation.
     * @param :uint8_t *polyPara: input irreducible polynomial key buffer, 16 bytes size,
     *                        the lowest bit shall be 1, if not, force to 1.
     */
    typedef struct _Mt_Para_Type
    {
        uint32_t hashLen;
        uint8_t *hashKey;
        uint32_t polyLen;
        uint8_t *polyPara;
    } Mt_Para_Type_s;

    /***
     * @description: get number of hash devices
     * @return : <=0: no device available, >0: n devices available
     */
    int mthash_pcie_get_count(void);

    /***
     * @description: Open and return an available hash PCIe device.
     * @param :int  devid: -1: Open the device file automatically assigned by the system,
     * 					   0~mthash_pcie_dev_init return value: open the device specified by user input.
     * @return : -1:failed; >=0:succeed device id
     */
    int mthash_pcie_open(int devid);

    /***
     * @description: device deinit
     * @return :
     */
    void mthash_pcie_close(int devid);

    /***
     * @description: Hash calculation function interface
     * @param :int  devid: device id
     * @param :Mt_Data_Type_s *src: input data, input data to be hashed and verified
     * @param :Mt_Para_Type_s *tHashParam: input data parameters
     * @param :Mt_Data_Type_s *dst: output, calculated result
     * @return :0:succeed, -1:failed
     */
    int mthash_pcie_get_hashvalue(int devid, Mt_Data_Type_s *src, Mt_Para_Type_s *tHashParam, Mt_Data_Type_s *dst);

    /***
     * @description: hash device status query interface
     * @param :int  devid: device id
     * @param :int  *dst: the value of PCIe device, the actual temperature value is dst / 100
     * @return : 0: device working normally,
     * 			 other: Byte0:bit0: PCIe device work status, 0:ok, 1: error, if this bit error, the device shall be reset.
     * 				    Byte0:bit1: PCIe device temperature, 0:ok, 1: error
     *  			    Byte0:bit3~7: default 0, reserved
     */
    int mthash_pcie_get_devstatus(int devid, int *dst);

    /***
     * @description: reset pcie device status,The internal detection is carried out in a cycle of 5 seconds.
     *               In case of abnormality, the API will automatically software reset the device.
     * @param :int  devid: device id
     * @return : 0: success, -1: failed
     */
    int mthash_pcie_dev_reset(int devid);

    /***
     * @description: get hash pcie API version
     * @param :char  *libinfo: this needs 6 bytes of memory space
     * @return :
     */
    void mthash_pcie_get_libraryinfo(char *libinfo);

    /***
     * @description: get hash pcie driver version
     * @param :char  *driverinfo: this needs 6 bytes of memory space
     * @return :
     */
    void mthash_pcie_get_driverinfo(char *driverinfo);

#ifdef __cplusplus
}
#endif
#endif
