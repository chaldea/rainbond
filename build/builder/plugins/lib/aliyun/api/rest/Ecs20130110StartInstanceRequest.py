'''
Created by auto_sdk on 2015.04.21
'''
from aliyun.api.base import RestApi
class Ecs20130110StartInstanceRequest(RestApi):
	def __init__(self,domain='ecs.aliyuncs.com',port=80):
		RestApi.__init__(self,domain, port)
		self.InstanceId = None

	def getapiname(self):
		return 'ecs.aliyuncs.com.StartInstance.2013-01-10'
