//logs.js
var util = require('../../utils/util.js')
Page({
  data: {
    ret: []
  },
  onLoad: function () {
    var that = this;
    wx.request({
      url: 'http://182.92.200.3:8080/api/v1/bless',
      data: {
        x: '' ,
        y: ''
      },
      header: {
          'content-type': 'application/json'
      },
      success: function(res) {
          that.setData({
              ret: res.data
          })
      },
    });
  }
})
