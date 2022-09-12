import COS from 'cos-js-sdk-v5';
import DateFormat from '@xfl/date-format/DateFormat.js';

var dateformat = new DateFormat();

const Bucket = 's1-1309644651'; //存储桶的名称，命名规则为 BucketName-APPID，此处填写的存储桶名称必须为此格式
const Region = 'ap-shanghai'; //存储桶所在地域
//创建一个 COS SDK 实例
// SECRETID 和 SECRETKEY请登录 https://console.cloud.tencent.com/cam/capi 进行查看和管理
const cos = new COS({
  SecretId: 'AKIDVptgDUTCepjLhncAVaOvBG6fxyiBNIxL',
  SecretKey: '8zauxiCQNPAftqWu2tSkmWzB7mAI8i13',
});

//创建存储桶
// cos.putBucket(
//   {
//     Bucket: Bucket,
//     Region: Region,
//   },
//   function (err, data) {
//     console.log(err || data);
//   }
// );

// cos.getBucket({
//     Bucket: Bucket, /* 必须 */
//     Region: Region,     /* 存储桶所在地域，必须字段 */
//     Prefix: 'web/',           /* 非必须 */
// }, function(err, data) {
//     console.log(err || data.Contents);
// });

//删除
const deleteFileToTencentClound = (Key) => {
  return new Promise((resolve, _) => {
    cos.deleteObject(
      {
        Bucket: Bucket,
        Region: Region,
        Key: Key,
      },
      function (err, data) {
        console.log(err || data);
        resolve(data);
      },
    );
  });
};

//上传图片到腾讯云
const uploadFileToTencentClound = (filename, file) => {
  return new Promise((resolve, reject) => {
    let Key = `web/${dateformat.toString('yyyy0m')}/${dateformat.timestamp}${filename}`;
    cos.putObject(
      {
        Bucket: Bucket,
        Region: Region,
        Key: Key,
        Body: file,
        onProgress: function (info) {
          console.log('[cos.postObject-info]', info);
        },
      },
      function (err, data) {
        console.log('[cos.postObject-success]', err || data);
        if (data && data.Location && data.statusCode == 200) {
          resolve({ location: `https://${data.Location}`, Key: Key });
        }
        if (err) {
          reject(err);
        }
      },
    );
  });
};

export { uploadFileToTencentClound, deleteFileToTencentClound };
