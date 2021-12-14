# Chapter 8 异常处理

## 8.1 error

&emsp;&emsp;Ge语言追求极简的设计哲学在error特性上体现得淋澳尽致。

&emsp;&emsp;在Go 1.13之前长达10余年的时间里，标准库对error的支持非常有限，仅有errors.New()
和fmt.Errorf()两个函数用来构造error实例。然而Go语言仅提供了eror的内置接口定义
(type error interface)，这样开发者可以定义任何类型的error，并可以存储任意的内容。

&emsp;&emsp;在Go 1.13之前，已经有很多开源项目试图扩展标准库的error以满足实际项目的需求，比
如pkg/errors,该项目被大量应用于诸如Kubernetes这样的大型项目中。如果读者之前了解过
这类项目，可能会发现Go1.13针对error的优化与这些开源项目有异曲同工之妙。

&emsp;&emsp;Go1.13在保持对原有error兼容的前提下，提供了新的error类型，新的error类型在函数
间传递时可以保存原始的error信息，本节我们称这类error为链式error。

&emsp;&emsp;本节先从基础的error讲起，包括error的基本用法、痛点，接着延伸到Go 1.13的改进，
方便读者了解error的演进脉络。

### 8.1.2 基础error

本节聚焦Go 1.13之前error的设计和使用，包括以下内容：
- error的类型；
- 如何创建error
- 如何自定义error;
- 如何检查error；
- 如何处理error；
- 使用error的痛点；

1.error接口  
2.创建error