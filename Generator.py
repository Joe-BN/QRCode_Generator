import qrcode
 qr = qrcode.QRCode(
     version = 1,
     error_correction = qrcode.constant.ERROR_CORRECT_L,
     box_size = 20,
     border = 20
 )
qr.add_data("_____") #add the link or the data that you want here.
qr.make(fit = True)
img = qr.make_image(fill_color = "black", back_color = "white")
img.save("______.png") #add the name of the Pic File that the qrcode will be saved in.
